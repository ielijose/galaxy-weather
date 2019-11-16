import React from 'react';

import {
  XYPlot,
  XAxis,
  YAxis,
  HorizontalGridLines,
  LineSeries,
  MarkSeries,
} from 'react-vis';

const sun = [{ x: 0, y: 0, size: 1 }];

export default function Galaxy({ data }) {
  return (
    <XYPlot
      xDomain={[-200, 200]}
      yDomain={[-200, 200]}
      width={400}
      height={400}
    >
      <HorizontalGridLines />
      <LineSeries color="red" opacity={0.1} data={data} />

      <MarkSeries data={sun} />

      <MarkSeries data={data} />

      <XAxis />
      <YAxis />
    </XYPlot>
  );
}
