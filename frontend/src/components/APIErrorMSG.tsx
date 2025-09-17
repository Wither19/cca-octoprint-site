type APIErrorMSGProps = {
  msg?: string;
}

function APIErrorMSG({ msg = "Could not reach printer API"}: APIErrorMSGProps) {
  return <div className="display-6 text-center">{msg}</div>
}

export default APIErrorMSG;