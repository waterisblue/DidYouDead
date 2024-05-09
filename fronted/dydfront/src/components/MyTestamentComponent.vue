<template>
    <v-data-table-virtual v-model:expanded="expanded" :headers="dessertHeaders" :items="desserts" item-value="TestamentFileName" show-expand>
        <template v-slot:top>
            <v-toolbar flat>
                <v-toolbar-title>遗嘱管理</v-toolbar-title>
            </v-toolbar>
        </template>
        <template v-slot:item.TestamentFileName="{ item }">  
            <a :href="'http://localhost:8888/static/testament/' + item.TestamentFileName">{{ item.TestamentFileName }}</a>
        </template>
        <template v-slot:item.IsActive="{ item }">  
            <v-checkbox v-model="item.IsActive"></v-checkbox>
        </template>
 
        <template v-slot:expanded-row="{ columns, item }">
            <tr>
                <td :colspan="columns.length" style="text-indent: 2rem;">
                    {{ item.TestamentDetail }}
                </td>
            </tr>
        </template>
    </v-data-table-virtual>
</template>
<script setup>
import axios from '../axios';
import { ref } from 'vue';

let expanded = ref([])
let dessertHeaders = ref([
    {
        title: '遗嘱名',
        align: 'start',
        sortable: false,
        key: 'TestamentName',
    },
    { title: '遗嘱形式', key: 'TestamentStyle' },
    { title: '遗嘱创建时间', key: 'CreateDate' },
    { title: '遗嘱过期时间', key: 'ExpireDate' },
    { title: '遗嘱文件', key: 'TestamentFileName' },
    { title: '是否激活', key: 'IsActive' },
    { title: '', key: 'data-table-expand' },
])
let desserts = ref()
axios.get('/loginafter/gettestament').then(res => {
    desserts.value = res.data.data.testaments
})
</script>