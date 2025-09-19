import "./App.scss"

import { useEffect, useState } from "react"
import axios from "axios"

import type { PrinterStatus } from "@jamesgopsill/octoprint-client"

import { Badge } from "react-bootstrap"

import PrinterLink from "./components/PrinterLink"
import TemperatureList from "./components/TemperatureList"
import APIErrorMSG from "./components/APIErrorMSG"

import { printerStateColors } from "./functions"

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]

  const [currentPrinter, setPrinter] = useState("a")
  const [printerState, setPrinterState] = useState<PrinterStatus>()

  var printerStateColor = printerStateColors(printerState?.state.flags)

  const getPrinterResponse = () => { axios.post("/api/thing").then((r) => setPrinterState(r.data)) }

  useEffect(getPrinterResponse, [currentPrinter])

  return (
    <>
      <div className="display-6 text-center mx-2">CCA 3D Printing</div>
      <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
        {printerList.map((p) => <PrinterLink key={p} printerID={p} click={() => setPrinter(p)} />)}
      </div>

      {printerState ? (
        <>
          <div className="display-6 text-center text-capitalize">
            Printer {currentPrinter}
            <Badge className="mx-3" bg={printerStateColor}>{printerState.state.text}</Badge>
          </div>
          <div id="temperature-container" className="mx-4">
            <TemperatureList temperatures={printerState!.temperature} />
          </div>
        </>
      ) : (
        <APIErrorMSG />
      )}
    </>
  )
}

export default App
