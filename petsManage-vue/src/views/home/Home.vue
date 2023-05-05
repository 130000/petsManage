<template>
    <el-row class="home" :gutter="20">
        <el-col :span="8" style="margin-top: 10px">
            <el-card shadow="hover">
                <div class="user">
                    <img src="../../assets/R-C.jfif" alt="" />
                    <div class="userInfo">
                        <p class="account">admin</p>
                        <p class="role">管理员</p>
                    </div>
                </div>

            </el-card>
            <div class="block text-center">
                <el-carousel shadow="hover" style="margin-top: 10px" height="320px">
                    <el-carousel-item v-for="item in petsPhoto" :key="item">
                        <el-image :src="item" alt="" :fit="cover"></el-image>
                    </el-carousel-item>
                </el-carousel>
            </div>
        </el-col>
        <el-col :span="16" style="margin-top: 10px">
            <el-card style="height: 280px;">
                <div ref="echart" style="height: 280px;">
                </div>
            </el-card>
            <div class="graph">
                <el-card style="height: 260px;">
                    <div ref="userechart" style="height: 240px"></div>
                </el-card>
                <el-card style="height: 260px;">
                    <div ref="videoechart" style="height: 240px"></div>
                </el-card>
            </div>
        </el-col>
</el-row>
</template>
<script setup>
import { onMounted, ref, getCurrentInstance } from 'vue';
import * as echarts from 'echarts'
const { proxy } = getCurrentInstance()
onMounted(() => {
    getBarChart()
    getPieChart()
    getPetsPhoto()
})
const petsPhoto = ref([])
const getPetsPhoto = async () => {
    let res = await proxy.$api.getPetsPhoto()
    petsPhoto.value = res
}
const getBarChart = () => {
    const myChart3 = echarts.init(proxy.$refs.echart)
    myChart3.setOption(barChartOption)
}
const getPieChart = () => {
    const myChart = echarts.init(proxy.$refs.userechart)
    const myChart2 = echarts.init(proxy.$refs.videoechart)
    myChart.setOption(pieChartOption)
    myChart2.setOption(pieChartOption2)
}
const barChartOption = {
    title: {
        text: 'xx社区宠物种类汇总',
        left: 'center'
    },
    xAxis: {
        data: [
            '狗',
            '猫',
            '其他'
        ]
    },
    yAxis: {},
    series: [
        {
            type: 'bar',
            data: [9, 2, 1]
        }
    ]
}

const pieChartOption2 = {
    title: {
        text: 'xx社区宠物不良记录情况',
        left: 'center'
    },
    series: [
        {
            type: 'pie',
            data: [
                {
                    value: 35,
                    name: '有不良记录'
                },
                {
                    value: 50,
                    name: '无不良记录'
                }
            ]
        }
    ]
}
const pieChartOption = {
    title: {
        text: 'xx社区宠物疫苗接种情况',
        left: 'center'
    },
    series: [
        {
            type: 'pie',
            roseType: 'area',
            data: [
                {
                    value: 35,
                    name: '接种1次'
                },
                {
                    value: 50,
                    name: '接种2次及以上'
                },
                {
                    value: 15,
                    name: '未接种疫苗'
                }
            ]
        }
    ]
}
</script>
<style lang="less" scoped>
.home {
    .user {
        display: flex;
        align-items: center;
        padding-bottom: 20px;
        border-bottom: 1px solid #ccc;
        margin-bottom: 20px;

        img {
            width: 150px;
            height: 150px;
            border-radius: 50%;
            margin-right: 40px;
        }
    }

    .login-info {
        p {
            line-height: 30px;
            font-size: 14px;
            color: #999;

            span {
                color: #666;
                margin-left: 70px;
            }
        }
    }

    .num {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;

        .el-card {
            width: 32%;
            margin-bottom: 20px;
        }

        .icons {
            width: 80px;
            height: 80px;
            font-size: 30px;
            text-align: center;
            line-height: 80px;
            color: #fff;
        }

        .details {
            margin-left: 15px;
            display: flex;
            flex-direction: column;
            justify-content: center;

            .num {
                font-size: 30px;
                margin-bottom: 10px;
            }

            .txt {
                font-size: 14px;
                text-align: center;
                color: #999;
            }
        }
    }

    .graph {
        margin-top: 20px;
        display: flex;
        justify-content: space-between;

        .el-card {
            width: 48%
        }
    }
}
</style>