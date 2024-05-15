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
        <v-select v-model="sex" :items="['男', '女']" :rules="[v => !!v || '必须输入性别']" label="性别" required></v-select>
      </div>
      <div class="nameInputs">
        <v-menu v-model="menuVisible" :close-on-content-click="false" transition="scale-transition" offset-y
          max-width="290px" min-width="290px">
          <template v-slot:activator="{ props }">
            <v-text-field v-model="orderTime" label="预约时间" prepend-icon="mdi-calendar" readonly
              v-bind="props"></v-text-field>
          </template>
          <v-date-picker v-model="selectTime" @update:modelValue="checkTime"></v-date-picker>
        </v-menu>
      </div>
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
let idcard = ref('')
let phone = ref('')
let locate = ref('')
let sex = ref('')
let orderTime = ref('')
let checkbox = ref(false)
let menuVisible = ref(false)
let selectTime = ref(new Date())

const router = useRouter()

import { useTimeLineStore, useFireServiceChoose } from '@/stores/user';
const { setTimeLineInfo } = useTimeLineStore()

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
  service.OrderTime = selectTime.value
  setService(service)
  router.push(path)
}

function checkTime(){
  menuVisible.value = false
  let dateFormat = selectTime.value
  orderTime.value = `${dateFormat.getFullYear()}-${dateFormat.getMonth() + 1}-${dateFormat.getDate()}`
  
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
      width: 83vh;
    }

    .houseInputs {
      width: 177vh;
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
