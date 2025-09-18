import type { TemperatureObject } from "../types";

import TemperatureListItem from "./TemperatureListItem";

type TemperatureListProps = {
  temperatures: TemperatureObject;
};

function TemperatureList({ temperatures }: TemperatureListProps) {
  return (
    <>
      <TemperatureListItem name="Bed" temperature={temperatures["bed"]} />
      <TemperatureListItem name="Tool-0" temperature={temperatures["tool0"]} />
      <TemperatureListItem name="Tool-1" temperature={temperatures["tool1"]} />
      <TemperatureListItem name="Tool-2" temperature={temperatures["tool2"]} />
      <TemperatureListItem name="Tool-3" temperature={temperatures["tool3"]} />
      <TemperatureListItem name="Tool-4" temperature={temperatures["tool4"]} />
    </>
  )
}

export default TemperatureList;