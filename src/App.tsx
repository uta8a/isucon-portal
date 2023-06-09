import { useState, useEffect } from "react";
import viteLogo from "/vite.svg";
import "./App.css";
import variables from "../variables.json";

import { Graph } from "./Graph";
import {useData} from "./useData";

function App() {
  useEffect(() => {
    document.title = variables.title;
  }, []);
  const data = useData(variables.url);
  console.log(data);

  return (
    <>
      <h1>{variables.title}</h1>
      <Graph />
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </>
  );
}

export default App;
