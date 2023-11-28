<script setup>

import ItemRow from "./ItemRow.vue";
import {ref, watch} from "vue";
import {getItemDataList, getItemDataListByCompany} from "@/api/itemData";

const props = defineProps({
    companyName: String
})

const dataList1 = ref([])
const groupCountNumber = ref(0)
const groupCount = ref(0)

watch(() => props.companyName, (newName) => {
    getItemDataListByCompany({company: newName}).then((res) => {
        dataList1.value = res.data.list
        console.log(dataList1.value.length)
        groupCountNumber.value = dataList1.value.length / 3
        groupCount.value = Number.isInteger(groupCountNumber.value) ? groupCountNumber.value : Math.floor(groupCountNumber.value) + 1
    })
})


</script>

<template>
    <div>
        <template v-for="i in [...Array(groupCount).keys()]">
            <ItemRow :slice-numbers="[groupCount, i * 3, (i + 1) * 3 < dataList1.length ? (i + 1) * 3 : dataList1.length]" :data-list="dataList1.slice(i * 3, (i + 1) * 3 < dataList1.length ? (i + 1) * 3 : dataList1.length)" />
            <a-divider />
        </template>
    </div>

</template>

<style scoped>

.ant-divider-horizontal {
    margin: 0;
}
</style>