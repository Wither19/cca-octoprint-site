export type TemperatureData = {
  Actual: number;
  Target: number;
  Offset: number;
};

type NoHistoryTemperature = Omit<keyof PrinterState, TemperatureHistory[]>;

export type ModifiedTemperatureData = {
  Name: string;
  Actual: number;
  Target: number;
  Offset: number;
};

export type TemperatureHistory = {
  Time: number;
  Tool: { [key: string]: TemperatureData };
};

export type PrinterStateFlags = {
  Operational: boolean;
  Paused: boolean;
  Printing: boolean;
  Cancelling: boolean;
  Pausing: boolean;
  SDReady: boolean;
  Error: boolean;
  Ready: boolean;
  ClosedOrError: boolean;
};

export type PrinterState = {
  Flags: PrinterStateFlags;
  Text: string;
  Error: string;
  Temperature: {
    Bed: TemperatureData;
    Tool0: TemperatureData;
    Tool1: TemperatureData;
    Tool2: TemperatureData;
    Tool3: TemperatureData;
    Tool4: TemperatureData;
    History: TemperatureHistory[];
  };
};

export type ModifiedPrinterResponse = {
  PrinterName: string;
  State: PrinterState;
  StateColor: string;
  Temperature: ModifiedTemperatureData[];
  TemperatureHistory: TemperatureHistory[];
};