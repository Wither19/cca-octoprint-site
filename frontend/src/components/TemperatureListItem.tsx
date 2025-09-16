import type { ModifiedTemperatureData } from "../types";

function TemperatureListItem({ temperature }: { temperature: ModifiedTemperatureData }) {
  return (
    <div id={temperature.Name} className="mb-3">
      <h5 className="text-capitalize">{temperature.Name.replace("-", " ")}
      </h5>
      <div>Actual: {temperature.Actual} &deg;C   
      </div>
      <div>Target: {temperature.Target} &deg;C
      </div>
    </div>
  )
}

export default TemperatureListItem;