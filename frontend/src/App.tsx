import "./App.scss"

import { useEffect, useState } from "react"
import axios from "axios"

import type { PrinterState, TemperatureObject } from "./types"

import PrinterLink from "./components/PrinterLink"
import StatusBadge from "./components/StatusBadge"
import TemperatureListItem from "./components/TemperatureListItem"

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]

  const [currentPrinter, setPrinter] = useState("a")
  const [printerState, setPrinterState] = useState<PrinterState>()

  const [temp, setTemp] = useState<Omit<TemperatureObject, "History"> | null>(null);

  function getPrinterResponse() {
    axios.post("/api/thing").then((r) => setPrinterState(r.data))
  }

  function convertTemp() {
    if (printerState) {
      const newTemp = printerState!.Temperature;
      delete newTemp.History;

      setTemp(newTemp);
    } else {
      return;
    }
  }

  useEffect(getPrinterResponse, [currentPrinter])

  useEffect(convertTemp, [printerState])

  return (
    <>
      <div className="display-6 text-center mx-2">CCA 3D Printing</div>
      <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
        {printerList.map((printer) => (
          <PrinterLink key={printer} printerID={printer} click={() => setPrinter(printer)} />
        ))}
      </div>

      {printerState && temp ? (
        <>
          <div className="display-6 text-center text-capitalize">
            Printer {currentPrinter}
            <StatusBadge state={printerState} />
          </div>
          <div id="temperature-container" className="mx-4">
            {Object.keys(temp).map((k) => (
              <TemperatureListItem key={k} temperature={temp[k as keyof typeof temp]} name={k} />
            ))}
          </div>
        </>
      ) : (
        <div className="display-6 text-center">Could not reach Printer API</div>
      )}
    </>
  )
}

export default App
