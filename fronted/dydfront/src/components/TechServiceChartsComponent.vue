<template>
    <div class="charts">
        <div class="lineChart" ref="lineChartDom"></div>
        <div class="pieChart" ref="pieChartDom"></div>
        <div class="buyerPage" ref="buyerPage">
            <div class="title">配套心率测量设备</div>
            <div class="description">
                我们的智能手表不仅提供了先进的健康监测功能，还结合了时尚设计和便捷的使用体验。无论是关心健康数据、追求时尚潮流，还是需要方便的智能提醒，我们的智能手表都能满足你的需求。<br/>
                现在就选择我们的智能手表，让健康与时尚同行，成为你生活中不可或缺的一部分。<br/>
                如果你对我们的产品有任何疑问或想了解更多信息，请随时联系我们。感谢您的关注与支持！
            </div>
            <v-btn color="error" @click="toPay()">现在购买！</v-btn>
        </div>
    </div>
</template>

<script setup>
import * as echarts from 'echarts'
import { onMounted } from 'vue';
import { ref } from 'vue';
import axios from '../axios'

const lineChartDom = ref(null)
const pieChartDom = ref(null)

function toPay() {
    let backUrl = import.meta.env.VITE_BACK_URL

    window.open(`${backUrl}/pay?amount=298.00&subject="智能设备产品支付"`, '_blank');
}
onMounted(() => {
    let lineChartSet = echarts.init(lineChartDom.value)
    let pieChartSet = echarts.init(pieChartDom.value)

    // 心率随机数生成
    let base = +new Date(1968, 9, 3);
    let oneDay = 1000;
    let date = [];
    let data = [Math.random() * 80];
    for (let i = 1; i < 20000; i++) {
        var now = new Date((base += oneDay));
        date.push((now.getMonth() + 1) + '月' + (now.getDate() + 1) + '日' + (now.getHours() + 1) + '点' + (now.getMinutes() + 1) + '分');
        data.push(Math.floor(Math.random() * (120 - 70 + 1)) + 70);
    }

    let lineChartOption = {
        tooltip: {
            trigger: 'axis',
            position: function (pt) {
                return [pt[0], '10%'];
            }
        },
        title: {
            left: 'center',
            text: '心率实时变化'
        },
        toolbox: {
            feature: {
                dataZoom: {
                    yAxisIndex: 'none'
                },
                restore: {},
                saveAsImage: {}
            }
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: date
        },
        yAxis: {
            type: 'value',
            boundaryGap: [0, '100%']
        },
        dataZoom: [
            {
                type: 'inside',
                start: 0,
                end: 10
            },
            {
                start: 0,
                end: 10
            }
        ],
        series: [
            {
                name: 'Fake Data',
                type: 'line',
                symbol: 'none',
                sampling: 'lttb',
                itemStyle: {
                    color: 'rgb(255, 70, 131)'
                },
                areaStyle: {
                    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                        {
                            offset: 0,
                            color: 'rgb(255, 158, 68)'
                        },
                        {
                            offset: 1,
                            color: 'rgb(255, 70, 131)'
                        }
                    ])
                },
                data: data
            }
        ]
    };

    let pieChartOption = {
        tooltip: {
            trigger: 'item'
        },
        legend: {
            top: '5%',
            left: 'center'
        },
        series: [
            {
                name: '时间',
                type: 'pie',
                radius: ['40%', '70%'],
                avoidLabelOverlap: false,
                itemStyle: {
                    borderRadius: 10,
                    borderColor: '#fff',
                    borderWidth: 2
                },
                label: {
                    show: false,
                    position: 'center'
                },
                emphasis: {
                    label: {
                        show: true,
                        fontSize: 40,
                        fontWeight: 'bold'
                    }
                },
                labelLine: {
                    show: false
                },
                data: [
                    { value: 6, name: '昨日深睡眠时间' },
                    { value: 2, name: '昨日浅睡眠时间' },
                    { value: 16, name: '昨日清醒时间' }
                ]
            }
        ]
    };
    lineChartOption && lineChartSet.setOption(lineChartOption)
    pieChartOption && pieChartSet.setOption(pieChartOption)
}

)

</script>

<style scoped lang="less">
.charts {
    display: flex;
    flex-direction: row;
    height: 60vh;

    .lineChart {
        flex: 1;
    }

    .pieChart {
        flex: 1
    }

    .buyerPage {
        flex: 1;
        display: flex;
        flex-direction: column;
        border: solid 0.01rem rgb(112, 180, 193);
        border-radius: 1rem;
        background-color: rgb(203, 93, 93);
        padding: 4rem;
        .title {
            text-align: start;
            color: aliceblue;
            font-weight: 600;
            font-size: 2rem;
            flex: 1;
        }
        .description {
            flex: 3;
            color: rgb(222, 187, 187);
        }
    }
}
</style>