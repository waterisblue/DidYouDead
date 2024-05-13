<template>
  <v-container fluid class="outer">
    <div class="nameInputGroup">
      <div class="nameInputs">
        <v-text-field v-model="firstName" label="姓名"></v-text-field>
      </div>
      <div class="nameInputs">
        <v-text-field v-model="age" label="年龄"></v-text-field>
      </div>
      <div class="nameInputs">
        <v-select v-model="sex" :items="['男', '女']" :rules="[v => !!v || '必须输入性别']" label="性别"
          required></v-select>
      </div>
      <div class="nameInputs">
        <v-text-field v-model="ordertime" label="预约时间"></v-text-field>
      </div>
      <!-- <div class="nameInputs">
        <v-date-picker v-model="ordertime"></v-date-picker>
      </div> -->
    </div>
    <div class="nameInputGroup">
      <div class="nameInputs2">
        <v-text-field v-model="idcard" label="身份证号"></v-text-field>
      </div>
      <div class="nameInputs2">
        <v-text-field v-model="phone" label="手机号"></v-text-field>
      </div>
    </div>
    <div class="nameInputGroup">
      <div class="houseInputs">
        <v-text-field v-model="locate" label="家庭住址"></v-text-field>
      </div>
    </div>
    <div class="checkboxInput">
      <v-checkbox v-model="checkbox" label="Do you agree?" required></v-checkbox>
    </div>
    <div class="btnInputs">
      <v-btn class="btnInput" color="error" @click="checkUrl('/fireservice')">上一步</v-btn>
      <v-btn class="btnInput" color="info" @click="checkUrl('/fireservice/firemode')">下一步</v-btn>
    </div>

  </v-container>
</template>


<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

let firstName = ref('')
let age = ref('')
let orderTime = ref(null)
let idcard = ref('')
let phone = ref('')
let locate = ref('')
let sex = ref('')
let ordertime = ref('')
let checkbox = ref(false)

const router = useRouter()

import { useTimeLineStore, useFireServiceChoose } from '@/stores/user';
import { storeToRefs } from 'pinia';
const { timeLineStore, setTimeLineInfo } = useTimeLineStore()

const { timeLineInfo } = storeToRefs(timeLineStore)
const { getService, setService } = useFireServiceChoose()

let service = getService()

setTimeLineInfo('two')

function checkUrl(path) {
  service.Name = firstName.value
  service.Age = age.value
  service.OrderTime = orderTime.value
  service.Sex = sex.value
  service.IdNum = idcard.value
  service.Locate = locate.value
  service.PhoneNum = phone.value
  service.OrderTime = ordertime.value
  setService(service)
  router.push(path)
}
</script>

<style lang="less" scoped>
.outer {
  display: flex;
  flex-direction: column;
  justify-content: space-around;

  .nameInputGroup {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    height: 15vh;

    .nameInputs {
      width: 35vh;
    }

    .nameInputs2 {
      width: 78vh;
    }

    .houseInputs {
      width: 170vh;
    }
  }

  .checkboxInput {
    display: flex;
    justify-content: center;
  }

  .btnInputs {
    display: flex;
    justify-content: space-around;

    .btnInput {
      width: 20vh;
    }
  }
}
</style>
