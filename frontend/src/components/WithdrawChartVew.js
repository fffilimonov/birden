import React from 'react';
import {VictoryPie} from 'victory-pie';
import {ReactComponent as CoinIcon} from '../assets/images/coinSmall.svg';

export const WithdrawChartView = ({balance = 0, onClick, price}) => {

    return (
        <div style={{
            width: '280px',
            position: 'relative',
            background: '#060606',
            border: ' 1px solid #2D2D2F',
            borderRadius: '8px',
            paddingTop: 22,
            paddingLeft: 19,
            paddingRight: 19,
            boxSizing: 'border-box',
            height: 225,
            overflow: 'hidden'
        }}>
            <div style={{
                position: 'relative',
            }}>
                {/*pie bg*/}
                <div style={{position: 'absolute', inset: 0}}>
                    <VictoryPie
                        height={225}
                        width={225}
                        horizontal
                        padding={{top: 0, left: 0, right: 0, bottom: 0}}
                        padAngle={({datum}) => 0}
                        cornerRadius={({datum}) => 10}
                        innerRadius={225 * 0.42}
                        startAngle={90}
                        endAngle={-90}
                        animate={{duration: 250}}
                        colorScale={['#242426'].reverse()}
                        data={[{y: 1, label: ' '}].reverse()}
                    />
                </div>
                <VictoryPie
                    height={225}
                    width={225}
                    horizontal
                    padding={{top: 0, left: 0, right: 0, bottom: 0}}
                    padAngle={({datum}) => 0}
                    cornerRadius={({datum}) => 10}
                    innerRadius={225 * 0.42}
                    startAngle={90}
                    endAngle={-90}
                    animate={{duration: 250}}
                    colorScale={['#4B55D2', 'transparent'].reverse()}
                    data={[{y: balance, label: ' '}, {y: 1 - balance, label: ' '}].reverse()}
                />
                <div style={{
                    position: 'absolute', top: 80, left: 0, right: 0, fontWeight: 700,
                    fontSize: '48px',
                    lineHeight: '38px',
                    textAlign: 'center'
                }}>
                    {(Math.max(balance, 0)).toFixed(2)}
                </div>
                <div style={{
                    position: 'absolute',
                    top: 128,
                    left: -5,
                    right: -5,
                    display: 'flex',
                    justifyContent: 'space-between',
                    fontSize: 14,
                    lineHeight: '16px',
                    color: '#8A8E96'
                }}>
                    <span>0.00</span>
                    <span style={{display: 'flex', alignItems: 'center'}}><CoinIcon
                        style={{marginRight: 5}}/>{(price).toFixed(4)} CELO / SEC</span>
                    <span>1.00</span>
                </div>
            </div>
            <div onClick={onClick} style={{
                background: balance > 0 ? '#FFF' : '#4F4FD0',
                color: balance > 0 ? '#000' : '#fff',
                position: 'absolute',
                bottom: 0,
                left: 0,
                right: 0,
                height: 44,
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                cursor: 'pointer'
            }}>
                {balance <= 0 && (
                    <span>I want this stream</span>
                )}
                {balance > 0 && balance !== 1 && (
                    <>
                        <svg style={{marginRight: 17}} width="9" height="10" viewBox="0 0 9 10" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path
                                d="M1.35352 9.98633C1.57812 9.98633 1.7832 9.91309 2.05664 9.75684L7.91602 6.3584C8.36523 6.09961 8.58984 5.86035 8.58984 5.47949C8.58984 5.10352 8.36523 4.86426 7.91602 4.60059L2.05664 1.20215C1.7832 1.0459 1.57812 0.977539 1.35352 0.977539C0.894531 0.977539 0.523438 1.32422 0.523438 1.91016V9.04883C0.523438 9.63477 0.894531 9.98633 1.35352 9.98633Z"
                                fill="black"/>
                        </svg>
                        <span>Streaming</span>
                    </>
                )}
                {balance === 1 && (
                    <>
                        <svg style={{marginRight: 7}} width="12" height="12" viewBox="0 0 12 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path
                                d="M3.76855 11.0508C4.10059 11.0508 4.36426 10.9092 4.77441 10.5527L6.4248 9.10742H9.06152C10.5557 9.10742 11.4541 8.19434 11.4541 6.71484V2.84277C11.4541 1.3584 10.5557 0.450195 9.06152 0.450195H2.93848C1.44434 0.450195 0.545898 1.3584 0.545898 2.84277V6.71484C0.545898 8.19434 1.50293 9.10742 2.85547 9.10742H2.99219V10.1914C2.99219 10.7236 3.28516 11.0508 3.76855 11.0508ZM4.07129 9.6543V8.30176C4.07129 7.96973 3.91504 7.84766 3.61719 7.84766H3.0166C2.20117 7.84766 1.80566 7.4375 1.80566 6.63672V2.91602C1.80566 2.11523 2.20117 1.70996 3.0166 1.70996H8.9834C9.79395 1.70996 10.1943 2.11523 10.1943 2.91602V6.63672C10.1943 7.4375 9.79395 7.84766 8.9834 7.84766H6.33203C5.99512 7.84766 5.84375 7.91113 5.59961 8.15527L4.07129 9.6543ZM3.59277 4.28809C3.59277 4.87402 3.96387 5.30859 4.51562 5.30859C4.75977 5.30859 4.96973 5.24512 5.10645 5.08398H5.18945C5.05762 5.44531 4.70117 5.72852 4.27148 5.83105C4.12012 5.86523 4.05176 5.95312 4.05176 6.07031C4.05176 6.21191 4.16895 6.30469 4.32031 6.30469C4.88184 6.30469 5.80957 5.63086 5.80957 4.52246C5.80957 3.7998 5.37012 3.2334 4.66699 3.2334C4.04688 3.2334 3.59277 3.67285 3.59277 4.28809ZM6.26367 4.28809C6.26367 4.87402 6.63477 5.30859 7.18164 5.30859C7.42578 5.30859 7.64062 5.24512 7.77246 5.08398H7.85547C7.72852 5.44531 7.36719 5.72852 6.94238 5.83105C6.78613 5.86523 6.71777 5.95312 6.71777 6.07031C6.71777 6.21191 6.83496 6.30469 6.98633 6.30469C7.54785 6.30469 8.48047 5.63086 8.48047 4.52246C8.48047 3.7998 8.03613 3.2334 7.33301 3.2334C6.71777 3.2334 6.26367 3.67285 6.26367 4.28809Z"
                                fill="black"/>
                        </svg>
                        <span>Thanks a lot!</span>
                    </>
                )}
            </div>
        </div>
    )
}