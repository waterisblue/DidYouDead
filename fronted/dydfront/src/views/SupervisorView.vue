<template>
  <v-container fluid class="outScreen">
    <div class="toolBar">
      <v-btn icon class="backBtn" @click="checkUrl('/bubble')">&lt; </v-btn>
      <div class="rateText">心率检测</div>
    </div>

    <v-table class="tables">
      <thead>
        <tr>
          <th class="text-left">
            姓名
          </th>
          <th class="text-left">
            心率
          </th>
          <th class="text-left">
            时间
          </th>
          <th class="text-left">
            备注
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.Id">
          <td>{{ item.Account }}</td>
          <td>{{ item.HeartRate }}</td>
          <td>{{ item.CreateDate.replace('T', ' ').replace('Z', '') }}</td>
          <td> <v-btn :color="checkColor(item.HeartRate)">{{ checkRate(item.HeartRate)}} </v-btn></td>

        </tr>
      </tbody>
    </v-table>
  </v-container>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from '../axios';
import { useRouter } from 'vue-router';

const items = ref([])
const fetchData = () => {
  axios.post('/loginafter/getDetectionData')
    .then(response => {
      items.value = response.data.data;
    })
    .catch(error => {
      console.error('Error fetching data: ', error);
    });
};

const router = useRouter()

function checkUrl(path) {
  router.push(path)
}

function checkRate(rate) {
  if (rate == 0) return "已死亡"
  if (rate < 50) return "衰竭"
  return "正常"
}

function checkColor(rate) {
  if (rate == 0) return "error"
  if (rate < 50) return "warning"
  return "info"
}
onMounted(() => {
  fetchData(); // 初始加载数据

  // 每3秒请求一次数据
  // TODO: 退出时clearInterval
  setInterval(() => {
    fetchData();
  }, 3000);
});
</script>

<style scoped lang="less">
.outScreen {
  background-color: aliceblue;
  min-height: 100vh;
  display: flex;
  flex-direction: column;

  .toolBar {
    flex: 1;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: start;
  }
  .tables {
    flex: 3;
    background-color: rgba(130, 170, 205, 0.173);
  }

  .rateText {
    font-size: large;
    margin-left: 1rem;
  }

  .backBtn {
    font-size: 1.5rem;
    color: rgb(133, 154, 173);
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>