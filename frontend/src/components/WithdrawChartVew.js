import React, {memo} from 'react';
import {VictoryPie} from "victory-pie";

export const WithdrawChartView = memo((props) => {
    const height = 225
    const value = 100
    return (
        <div style={{height: 225, width: '100%'}}>
            <VictoryPie
                height={height}
                width={height}
                horizontal
                padding={{top: 0, left: 0, right: 0, bottom: 0}}
                padAngle={({ datum }) => 2}
                cornerRadius={({ datum }) => 10}
                innerRadius={height * 0.42}
                startAngle={90}
                endAngle={-90}
                animate={{duration: 250}}
                colorScale={['linear-gradient(88.59deg, #4B55D2 0%, #5643CC 45.83%, #663FD3 100%)', 'transparent'].reverse()}
                data={[{ y: value, label: ' '}, {y: 1 - value, label: ' '}].reverse()}
            />
        </div>
    )
});