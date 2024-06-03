<template>
    <v-container class="outerTestament">
        <div class="btns">
            <v-btn class="btn" color="info" :disabled=editBtnDisabled @click="editBtn">编辑</v-btn>
            <v-btn class="btn" color="primary" variant="tonal" @click="saveBtn">保存</v-btn>
        </div>
        <v-data-table-virtual v-model:expanded="expanded" :headers="dessertHeaders" :items="desserts"
            item-value="TestamentFileName" show-expand>
            <template v-slot:top>
                <v-toolbar flat>
                    <v-toolbar-title>遗嘱管理</v-toolbar-title>
                </v-toolbar>
            </template>
            <template v-slot:item.TestamentFileName="{ item }">
                <a :href="staticENV + item.TestamentFileName">{{ item.TestamentFileName }}</a>
            </template>
            <template v-slot:item.IsActive="{ item }">
                <v-checkbox :disabled=checkBoxDisabled v-model="item.IsActive"></v-checkbox>
            </template>

            <template v-slot:expanded-row="{ columns, item }">
                <tr>
                    <td :colspan="columns.length" style="text-indent: 2rem;">
                        {{ item.TestamentDetail }}
                    </td>
                </tr>
            </template>
        </v-data-table-virtual>
        <SnackBarComponent :snackText="snackText" :snackbar="snackBar" snackColor="waring" />
    </v-container>

</template>
<script setup>
import axios from '../axios';
import { ref } from 'vue';
import SnackBarComponent from './SnackBarComponent.vue';

// snackBar
let snackText = ref('')
let snackBar = ref(false)

let staticENV = ref(import.meta.env.VITE_STATIC_URL + '/testament/')
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

let checkBoxDisabled = ref(true)
let editBtnDisabled = ref(false)
function editBtn(){
    checkBoxDisabled.value = false
    editBtnDisabled.value = true
}

function saveBtn() {
    axios.post('/loginafter/updatetestamentactive', desserts.value).then(res => {
        if(res.data.code != 200){
            snackBar.value = true
            snackText.value = res.data.msg
            return
        }
        snackBar.value = true
        snackText.value = res.data.msg
    })
    checkBoxDisabled.value = true
    editBtnDisabled.value = false

}

</script>


<style lang="less" scoped>
.outerTestament {
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 100vh;
}
.btns {
    display: flex;
    flex-direction: row;
    justify-content: right;
    .btn {
        margin: 1rem;
    }
}
</style>