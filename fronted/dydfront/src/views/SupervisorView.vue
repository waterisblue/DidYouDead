<template>
    <v-container>
      <v-row>
        <v-col>
          <v-list>
            <v-list-item v-for="item in items" :key="item.Id">
              <v-list-item-content>
                <v-list-item-title>{{ item.Account }}</v-list-item-title>
                <v-list-item-title>{{ item.HeartRate }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
    </v-container>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  import axios from '../axios';
  
  const items = ref([]);
  
  const fetchData = () => {
    axios.post('/loginafter/getDetectionData')
      .then(response => {
        items.value = response.data.data;
      })
      .catch(error => {
        console.error('Error fetching data: ', error);
      });
  };
  
  onMounted(() => {
    fetchData(); // 初始加载数据
  
    // 每3秒请求一次数据
    setInterval(() => {
      fetchData();
    }, 3000);
  });
  </script>