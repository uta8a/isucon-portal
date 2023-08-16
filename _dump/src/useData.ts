import {
  useRef,
  useCallback,
  useSyncExternalStore,
} from "react";

// fetch data from benchmarker server
type Log = {
  timestamp: string;
  user: string;
  pass: boolean;
  score: number;
  success: number;
  fail: number;
  messages: string[];
};

type RenderedData = {
  logs: Log[];
  scoreboard: { [key: string]: number };
};

const fetchData = async (
  url: string,
  controller: AbortController
): Promise<RenderedData> => {
  const response = await fetch(url, { signal: controller.signal });
  const data = await response.text();
  const logs = data.split("\n");
  const logData: Log[] = [];
  const scoreboard: { [key: string]: number } = {};
  for (const rawLog of logs) {
    if (rawLog === "") continue;
    const log: Log = JSON.parse(rawLog);
    logData.push(log);
    scoreboard[log.user] = scoreboard[log.user]
      ? Math.max(log.score, scoreboard[log.user])
      : log.score; // maxを記録
  }
  const ret = {
    logs: logData,
    scoreboard: scoreboard,
  };
  return ret;
};

const useData = (url: string): RenderedData | undefined => {
  const data$ = useRef<RenderedData>();

  const subscribe = useCallback(
    (onStoreChange: () => void): (() => void) => {
      const controller = new AbortController();

      fetchData(url, controller).then((data) => {
        data$.current = data;

        onStoreChange();
      });

      return () => {
        controller.abort();
      };
    },
    [url]
  );

  return useSyncExternalStore(subscribe, () => data$.current);
};

export {fetchData, useData};
