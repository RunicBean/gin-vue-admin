<script setup>
import {watch, ref} from "vue";
import {getItemDataListByCompany} from "@/api/itemData";
import DetailedItems from "@/components/external/layoutOne/DetailedItems.vue";

const groupMap = ref({})
const isGroupMapUpdated = ref(false)
const props = defineProps({
    companyName: String
})
const mainItemsContainer = ref(null)

function groupByGroupTitle(dataList) {
    let groupMap = {}
    for (let dataObj of dataList) {
        if (Object.keys(groupMap).indexOf(dataObj.groupTitle) < 0) {
            groupMap[dataObj.groupTitle] = []
        }
        groupMap[dataObj.groupTitle].push(dataObj)
    }
    return groupMap
}

watch(() => props.companyName, (newName) => {
    getItemDataListByCompany({company: newName}).then((res) => {
        groupMap.value = groupByGroupTitle(res.data.list)
        isGroupMapUpdated.value = true


    })

})
</script>

<template>
    <div>
        <div v-if="isGroupMapUpdated"></div>
        <template v-for="(dataL, groupName) in groupMap">
            <h2>{{groupName}}</h2>
            <DetailedItems :data-l="dataL" />
        </template>
    </div>
</template>

<style scoped lang="scss">
h2 {
    margin: 10px 12px;
}
</style>