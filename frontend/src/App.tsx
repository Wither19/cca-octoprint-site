import "./App.scss"

import { useEffect, useState } from "react"
import axios from "axios"

import type { ModifiedPrinterResponse, ModifiedTemperatureData } from "./types"

import PrinterLink from "./components/PrinterLink"
import StatusBadge from "./components/StatusBadge"
import TemperatureListItem from "./components/TemperatureListItem"

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]

  const [currentPrinter, setPrinter] = useState("a")
  const [printerData, setPrinterData] = useState<ModifiedPrinterResponse>()

  function getPrinterResponse() {
    axios.post("/api/thing").then((r) => setPrinterData(r.data))
  }


  useEffect(getPrinterResponse, [currentPrinter])

  return (
    <>
      <div className="display-6 text-center mx-2">CCA 3D Printing</div>
      <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
        {printerList.map((printer) => (
         <PrinterLink key={printer} printerID={printer} click={() => setPrinter(printer)} />
        ))}
      </div>

      {printerData ? (
        <>
          <div className="display-6 text-center text-capitalize">
            Printer {currentPrinter}
           <StatusBadge state={printerData.State} />
          </div>
          <div id="temperature-container" className="mx-4">
            {printerData && printerData.Temperature.map((temp: ModifiedTemperatureData) => (
              <TemperatureListItem temperature={temp} key={`temp-${temp.Name}`} />
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
