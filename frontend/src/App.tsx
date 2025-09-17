import "./App.scss"

import { useEffect, useState } from "react"
import axios from "axios"

import type { ModifiedTemperatureData, PrinterState, TemperatureData, TemperatureHistory, NoHistoryTemperature, TemperatureObject } from "./types"

import PrinterLink from "./components/PrinterLink"
import StatusBadge from "./components/StatusBadge"
import TemperatureListItem from "./components/TemperatureListItem"

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]

  const [currentPrinter, setPrinter] = useState("a")
  const [printerState, setPrinterState] = useState<PrinterState>()

  var tempatureList: ModifiedTemperatureData | null;

  function getPrinterResponse() {
    axios.post("/api/thing").then((r) => setPrinterState(r.data))
  }


  useEffect(getPrinterResponse, [currentPrinter])

  useEffect(() => {


    let tempList: Omit<TemperatureObject, "History"> = printerState!.Temperature.filter((t) => !Array.isArray(t));
  }, [printerState])

  return (
    <>
      <div className="display-6 text-center mx-2">CCA 3D Printing</div>
      <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
        {printerList.map((printer) => (
          <PrinterLink key={printer} printerID={printer} click={() => setPrinter(printer)} />
        ))}
      </div>

      {printerState ? (
        <>
          <div className="display-6 text-center text-capitalize">
            Printer {currentPrinter}
            <StatusBadge state={printerState} />
          </div>
          <div id="temperature-container" className="mx-4">

          </div>
        </>
      ) : (
        <div className="display-6 text-center">Could not reach Printer API</div>
      )}
    </>
  )
}

export default App
