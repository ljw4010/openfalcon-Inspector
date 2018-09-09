package handler

import (
	"fmt"
	"math"
	"models"
	"strings"
	"time"
)

func HandlerCheck() {
	checkPoints := ImportExcel()
	allScreen := GetALLscreen()
	var isokCount int = 0
	currentTime := time.Now().Unix()

	for _, point := range checkPoints {
		historyTime := currentTime - int64(point.SpanTime*3600)
		graphs := getPointScreen(point, allScreen)

		for _, g := range graphs {
			var reqdata = &models.ReqHistory{
				Step:      60,
				StartTime: historyTime,
				HostNames: g.Endpoints,
				EndTime:   currentTime,
				Counters:  g.Counters,
				ConsolFun: point.ComMode,
			}
			graphValues := GetGraphHistory(reqdata)
			badpoints := judgeStatus(point.JudgeSymbol, point.Threshold, graphValues[0].Values)

			if badpoints == nil {
				point.IsAbnormal = true
			} else {
				point.IsAbnormal = false
				isokCount = isokCount + 1
				point.Desc = fmt.Sprintf("bad point: %+v", badpoints)
			}
		}
	}

	ExportExcel(checkPoints)

	if isokCount != 0 {
		Sendmail(false)
	} else {
		Sendmail(true)
	}

}

func getPointScreen(point *models.CheckTable, allScreen []models.DashBordScreen) []models.GraphInfo {
	for _, s := range allScreen {
		if strings.ToLower(s.Name) == strings.ToLower(strings.Trim(point.ParScreen, " ")) {
			for _, cscreen := range allScreen {
				if cscreen.Pid == s.Id {
					if strings.ToLower(cscreen.Name) == strings.ToLower(strings.Trim(point.ChildScreen, " ")) {
						return GetGraphByScreenId(cscreen.Id)
					}
				}
			}

		}
	}

	return nil
}

func judgeStatus(judgeSymbol string, v float64, vs []models.V) []models.V {
	var badValue []models.V
	for _, vsi := range vs {
		//		timeLayout := "2006-01-02 15:04:05"
		//		t := time.Unix(vsi.Timestamp, 0).Format(timeLayout)
		//		loc, _ := time.LoadLocation("Asia/Shanghai")
		//		theTime, _ := time.ParseInLocation(timeLayout, t, loc)
		//      log.Println(theTime, vsi.Value, judgeSymbol, v)
		if !checkIsTriggered(vsi.Value, judgeSymbol, v) {
			badValue = append(badValue, vsi)
		}
	}
	return badValue
}

func checkIsTriggered(leftValue float64, operator string, rightValue float64) (isTriggered bool) {

	switch operator {
	case "=", "==":
		isTriggered = math.Abs(leftValue-rightValue) < 0.0001
	case "!=":
		isTriggered = math.Abs(leftValue-rightValue) > 0.0001
	case "<":
		isTriggered = leftValue < rightValue
	case "<=":
		isTriggered = leftValue <= rightValue
	case ">":
		isTriggered = leftValue > rightValue
	case ">=":
		isTriggered = leftValue >= rightValue
	}

	return
}
