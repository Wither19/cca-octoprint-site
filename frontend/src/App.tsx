import { useEffect, useState } from "react";
import "./App.scss"

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]

  const [currentPrinter, setPrinter] = useState("a")

  useEffect(() => {
    callHandlerJSON("/api/thing")
  }, [currentPrinter])

  return (
    <>
      <div className="display-6 text-center mx-2">CCA 3D Printing</div>
      <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
        {printerList.map((printer) => (
          <a className="text-capitalize fs-3" onClick={() => setPrinter(printer)} href={`printer/${printer}`}>Printer {printer}</a>
        ))}
      </div>

      <div className="mt-3 mx-4 position-relative">
        <a href="/">
          <button className="btn btn-secondary">&lt;</button>
        </a>
      </div>

      <h2 className="display-6 text-center text-capitalize">
        {currentPrinter}
        <span className="badge bg-{{ .StateColor }} fs-6 position-relative">{{ .State.Text }}</span>
      </h2>
      <div id="temperature-container" className="mx-4">
        {{ range .Temperature }}
        <div id="{{ .Name }}" className="mb-3">
          <h5 className="text-capitalize">{{ .Name }}</h5>
          <div>Actual: {{ .Actual }} &deg;C</div>
          <div>Target: {{ .Target }} &deg;C</div>
        </div>
        {{ end }}
      </div>
    </>
  )
}

export default App
