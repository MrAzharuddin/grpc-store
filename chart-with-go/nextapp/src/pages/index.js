import Head from "next/head";
import { ChartRequest } from "../proto/chart_pb";
import { ChartDataServiceClient } from "../proto/chart_grpc_web_pb";
import { useState, useEffect } from "react";
import ReactECharts from "echarts-for-react";

var client = new ChartDataServiceClient("http://localhost:8000");
export default function Home() {
  const [lineX, setLineX] = useState([]);
  const [lineY, setLineY] = useState([]);
  const [barX, setBarX] = useState([]);
  const [barY, setBarY] = useState([]);

  const getLineChart = () => {
    // console.log("Line Chart Called");
    var req = new ChartRequest();
    var stream = client.getLineChartData(req, {});
    var x = [];
    var y = [];
    stream.on("data", function (response) {
      // console.log(response.getX(), response.getY());
      // setLineX(() => (lineX.push(response.getX())));
      // setLineY(() => (lineY.push(response.getY())));
      let now = response.getX();
      let date = new Date(now).toLocaleTimeString().replace(/^\D*/, "");
      // console.log(
      //   now.toString().length,
      //   "-> Len of Date",
      //   date.toString().length,
      //   "-> Len of Date"
      // );
      // console.log(now, "-> Date", date, "-> Date");
      // console.log(date.toLocaleTimeString().replace(/^\D*/, ""));

      x.push(date);

      y.push(response.getY());
      setLineX(x);
      setLineY(y);
    });
  };

  const getBarChart = () => {
    // console.log("Line Chart Called");
    var req = new ChartRequest();
    var stream = client.getBarChartData(req, {});
    var x = [];
    var y = [];
    stream.on("data", function (response) {
      // console.log(response.getX(), response.getY());
      // setLineX(() => (lineX.push(response.getX())));
      // setLineY(() => (lineY.push(response.getY())));
      let now = response.getX();
      let date = new Date(now).toLocaleTimeString().replace(/^\D*/, "");
      x.push(date);
      y.push(response.getY());
      setBarX(x);
      setBarY(y);
    });
  };

  useEffect(() => {
    getLineChart();
    getBarChart();
    console.log("LineX Data :", lineX);
    console.log("LineY Data :", lineY);
    console.log("Bar X Data :", barX);
    console.log("Bar Y Data :", barY);
    // setInterval(() => {
    //   getLineChart();
    //   console.log(lineX, lineY);
    // }, 5000);
  });
  const options = {
    grid: { top: "25%", right: "25%", bottom: "25%", left: "25%" },
    xAxis: [
      {
        type: "category",
        data: lineX.length > 10 ? lineX.slice(-10) : lineX,
      },
      {
        type: "category",
        data: barX.length > 10 ? barX.slice(-10) : barX,
      },
    ],
    yAxis: {
      type: "value",
    },
    series: [
      {
        data: lineY.length > 10 ? lineY.slice(-10) : lineY,
        type: "line",
        smooth: true,
      },
      {
        data: barY.length > 10 ? barY.slice(-10) : barY,
        type: "bar",
        smooth: true,
      },
    ],
    tooltip: {
      trigger: "axis",
      feature: {
        saveAsImage: { show: true },
      },
    },
  };

  const lineoptions = {
    grid: { top: "25%", right: "25%", bottom: "25%", left: "25%" },
    xAxis: {
      type: "category",
      data: lineX.length > 10 ? lineX.slice(-10) : lineX,
    },
    yAxis: {
      type: "value",
    },
    series: [
      {
        data: lineY.length > 10 ? lineY.slice(-10) : lineY,
        type: "line",
        smooth: true,
      },
    ],
    tooltip: {
      trigger: "axis",
    },
  };

  const baroptions = {
    grid: { top: "25%", right: "25%", bottom: "25%", left: "25%" },
    xAxis: {
      type: "category",
      data: barX.length > 10 ? barX.slice(-10) : barX,
    },
    yAxis: {
      type: "value",
    },
    series: [
      {
        data: barY.length > 10 ? barY.slice(-10) : barY,
        type: "bar",
        smooth: true,
        animationType: "scale",
        animationEasing: "elasticOut",
        animationDelay: function (idx) {
          return Math.random() * 200;
        },
      },
    ],
    tooltip: {
      trigger: "axis",
    },
  };

  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <div>
          <ReactECharts option={lineoptions} />
          <ReactECharts option={options} />
          <ReactECharts option={baroptions} />
        </div>
      </main>
    </>
  );
}
