import { Line } from "react-chartjs-2";
import variables from "../variables.json";
import {
  Chart as ChartJS,
  TimeScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  ChartOptions,
} from "chart.js";
import { useData } from "./useData";
import "chartjs-adapter-date-fns";
import { ja } from "date-fns/locale";

ChartJS.register(
  TimeScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

const options: ChartOptions<"line"> = {
  scales: {
    x: {
      type: "time",
      time: {
        unit: "hour",
      },
      adapters: {
        date: {
          locale: ja,
        },
      },
    },
  },
};

export const Graph = () => {
  const rawData = useData(variables.url);
  console.log(rawData);
  const labels = new Set();
  const mappedData = rawData?.logs.map((item) => {
    labels.add(item.user);
    return {
      x: item.timestamp,
      y: item.score,
      label: item.user,
    };
  });
  const data: { [key: string]: {x: string, y: number}[] }  = {}
  for (const item of mappedData ?? []) {
    if (data[item.label] === undefined) {
      data[item.label] = []
    }
    data[item.label].push({
      x: item.x,
      y: item.y,
    })
  }
  const values = []
  for (const key of labels) {
    values.push({
      label: key as string,
      borderColor: `#${Math.floor(Math.random()*16777215).toString(16)}`,
      data: data[key as string],
    })
  }
  const dataset = {
    datasets: values,
  };
  return rawData === undefined ? (
    <div>no image</div>
  ) : (
    <Line options={options} data={dataset} />
  );
};
