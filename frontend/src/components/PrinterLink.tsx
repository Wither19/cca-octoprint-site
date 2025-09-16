type PrinterLinkProps = { printerID: string, click: () => void }

function PrinterLink({ printerID, click }: PrinterLinkProps) {
  return  (
  <a style={{ cursor: "pointer" }} className="text-capitalize fs-3 text-decoration-none" onClick={click}>
    Printer {printerID}
  </a>
    )
}

export default PrinterLink