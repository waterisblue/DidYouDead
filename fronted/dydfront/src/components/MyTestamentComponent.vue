<template>
    <v-data-table-virtual v-model:expanded="expanded" :headers="dessertHeaders" :items="desserts" item-value="name" show-expand>
        <template v-slot:top>
            <v-toolbar flat>
                <v-toolbar-title>遗嘱管理</v-toolbar-title>
            </v-toolbar>
        </template>
        <template v-slot:item.isActive="{ item }">
            <v-checkbox v-model="item.isActive"></v-checkbox>
        </template>
        <template v-slot:expanded-row="{ columns, item }">
            <tr>
                <td :colspan="columns.length">
                    More info about {{ item.testamentContent }}
                </td>
            </tr>
        </template>
    </v-data-table-virtual>
</template>
<script setup>
import { ref } from 'vue';

let expanded = ref([])
let dessertHeaders = ref([
    {
        title: '遗嘱名',
        align: 'start',
        sortable: false,
        key: 'name',
    },
    { title: '遗嘱形式', key: 'testamentStyle' },
    { title: '遗嘱创建时间', key: 'createDate' },
    { title: '遗嘱过期时间', key: 'expireDate' },
    { title: '是否激活', key: 'isActive' },
    { title: '', key: 'data-table-expand' },
])
let desserts = ref([
    {
        name: '3月31日立',
        testamentStyle: '口述',
        createDate: '2023-03-31',
        expireDate: '2026-03-31',
        isActive: true,
        testamentContent: '撒大苏打撒旦撒大大十大撒大苏打撒旦2.撒大苏打撒旦撒大苏打撒旦哇撒大撒大撒v点三公分DVD法国二点产生的v奋斗吧'
    },
    {
        name: '4月28日补',
        testamentStyle: '音频',
        createDate: '2023-04-28',
        expireDate: '2026-04-28',
        isActive: true,
        testamentContent: '撒大苏打撒旦撒大大十大撒大苏打撒旦2.撒大苏打撒旦撒大苏打撒旦哇撒大撒大撒v点三公分DVD法国二点产生的v奋斗吧'
    },
    {
        name: '5月21日立',
        testamentStyle: '口述',
        createDate: '2023-05-21',
        expireDate: '2026-05-21',
        isActive: false,
        testamentContent: '撒大苏打撒旦撒大大十大撒大苏打撒旦2.撒大苏打撒旦撒大苏打撒旦哇撒大撒大撒v点三公分DVD法国二点产生的v奋斗吧'
    },
    {
        name: '6月21日立',
        testamentStyle: '视频',
        createDate: '2023-06-21',
        expireDate: '2026-06-21',
        isActive: false,
        testamentContent: '撒大苏打撒旦撒大大十大撒大苏打撒旦2.撒大苏打撒旦撒大苏打撒旦哇撒大撒大撒v点三公分DVD法国二点产生的v奋斗吧'
    },
])
</script>