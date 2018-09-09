package handler

import (
	"encoding/json"
	"fmt"
	"httpclient"
	"log"
	"models"

	"g"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func GetGraphByScreenId(id int64) []models.GraphInfo {
	var getGraphByScreenIdUrl = fmt.Sprintf("http://%s/api/v1/dashboard/graphs/screen/%%d", g.Config().ApiAddr)
	getGraphByScreenIdUrl = fmt.Sprintf(getGraphByScreenIdUrl, id)
	respBody := httpclient.Client("GET", getGraphByScreenIdUrl, nil)
	var allDashSc = []models.GraphInfo{}
	err := json.Unmarshal(respBody, &allDashSc)
	if err != nil {
		log.Fatal("get graphs failed by screen id ,err:", err.Error())
	}
	return allDashSc

}

func GetALLscreen() []models.DashBordScreen {
	var allDashboardScreenUrl = fmt.Sprintf("http://%s/api/v1/dashboard/screens", g.Config().ApiAddr)
	respBody := httpclient.Client("GET", allDashboardScreenUrl, nil)

	var allDashSc = []models.DashBordScreen{}
	err := json.Unmarshal(respBody, &allDashSc)
	if err != nil {
		log.Fatal("get all screen failed ,err:", err.Error())
	}
	return allDashSc

}

func GetGraphHistory(his *models.ReqHistory) []models.Resp {
	var historyUrl = fmt.Sprintf("http://%s/api/v1/graph/history", g.Config().ApiAddr)
	datslice, _ := json.Marshal(his)
	respBody := httpclient.Client("POST", historyUrl, datslice)

	res := []models.Resp{}
	err := json.Unmarshal(respBody, &res)
	if err != nil {
		fmt.Println("resp err:", err.Error())
	}
	//fmt.Println("res:", res)
	//	pts := make(plotter.XYs, len(res[0].Values))
	//	for i := range pts {
	//		pts[i].X = float64(res[0].Values[i].Timestamp)
	//		pts[i].Y = res[0].Values[i].Value
	//	}

	//	Plot(pts)
	return res
}

func PlotTest(points plotter.XYs) {

	p, _ := plot.New()
	p.Title.Text = "Hello Price"
	p.X.Label.Text = "Quantity Demand"
	p.Y.Label.Text = "Price"

	plotutil.AddLinePoints(p, points)

	p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}
