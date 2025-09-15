import { useEffect, useState } from "react";
import "./App.scss"
import type { ModifiedPrinterResponse, ModifiedTemperatureData } from "./types";
import axios from "axios";

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]

  const [currentPrinter, setPrinter] = useState("a")
  const [printerData, setPrinterData] = useState<ModifiedPrinterResponse | null>()

  function getPrinterResponse() {
    axios.post("/api/thing").then((r) => setPrinterData(r.data))
  }

  useEffect(getPrinterResponse, [currentPrinter])

  return (
    <>
      <div className="display-6 text-center mx-2">CCA 3D Printing</div>
      <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
        {printerList.map((printer) => (
          <a className="text-capitalize fs-3 cursor-pointer" onClick={() => setPrinter(printer)}>Printer {printer}</a>
        ))}
      </div>

      <div className="mt-3 mx-4 position-relative">
        <a href="/">
          <button className="btn btn-secondary">&lt;</button>
        </a>
      </div>

      {printerData ? (
        <>
      <h2 className="display-6 text-center text-capitalize">
        Printer {currentPrinter}
        <span className={`badge bg-${printerData.StateColor} fs-6 position-relative`}>{printerData.State.Text}</span>
      </h2>
      <div id="temperature-container" className="mx-4">
        {printerData && printerData.Temperature.map((temp: ModifiedTemperatureData) => (
          <div id={temp.Name} className="mb-3">
          <h5 className="text-capitalize">{temp.Name.replace("-", " ")}</h5>
          <div>Actual: {temp.Actual} &deg;C</div>
          <div>Target: {temp.Target} &deg;C</div>
        </div>
        ))}
      </div>
      </>
      ) : (
        <div className="display-3 text-center">Could not reach Printer API</div>
      )}
    </>
  )
}

export default App
