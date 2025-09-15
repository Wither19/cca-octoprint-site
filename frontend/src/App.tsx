import "./App.scss"

function App() {

  const printerList = ["a", "b", "c", "d", "e", "f"]


  return (
    <>
     <div className="display-6 text-center mx-2">CCA 3D Printing</div>
    <div className="mx-3 my-5 d-flex flex-wrap justify-content-around">
      {printerList.map((printer) => (
         <a className="text-capitalize fs-3" href={`printer/${printer}`}>Printer {printer}</a>
      ))}
    </div>
    </>
  )
}

export default App
