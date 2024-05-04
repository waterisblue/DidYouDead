<template>
  <v-container fluid class="outer">
    <div class="nameInputGroup">
      <div class="nameInputs">
        <v-text-field v-model="firstName" label="姓名"></v-text-field>
      </div>
      <div class="nameInputs">
        <v-text-field v-model="lastName" label="年龄"></v-text-field>
      </div>
      <div class="nameInputs">
        <v-text-field v-model="lastName" label="预约时间"></v-text-field>
      </div>
    </div>
    <div class="nameInputGroup">
      <div class="nameInputs2">
        <v-text-field v-model="firstName" label="身份证号"></v-text-field>
      </div>
      <div class="nameInputs2">
        <v-text-field v-model="lastName" label="手机号"></v-text-field>
      </div>
    </div>
    <div class="nameInputGroup">
      <div class="houseInputs">
        <v-text-field v-model="lastName" label="家庭住址"></v-text-field>
      </div>
    </div>
    <div class="checkboxInput">
      <v-checkbox v-model="checkbox" label="Do you agree?" required></v-checkbox>
    </div>
    <div class="btnInputs">
      <v-btn class="btnInput" color="error">上一步</v-btn>
      <v-btn class="btnInput" color="info">下一步</v-btn>
    </div>

  </v-container>
</template>


<script setup>
import { ref } from 'vue';

let firstName = ref('')
let lastName = ref('')
let checkbox = ref(false)
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
      width: 47vh;
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
