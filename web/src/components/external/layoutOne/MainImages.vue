<script setup>
import {onMounted, ref} from "vue";
import {getCarouselDataListByCompany} from "@/api/carouselData";

const props = defineProps({
    companyName: String
})

const dataList1 = ref([])
onMounted(() => {
    getCarouselDataListByCompany({company: props.companyName}).then((res) => {
        dataList1.value = res.data.list
    })
})
</script>

<template>
    <a-carousel autoplay>
        <template v-for="d in dataList1">
            <div>
                <img :src="d.imageLink" alt="">
            </div>
        </template>
    </a-carousel>
</template>

<style scoped>
img {
    width: 100%;
}
</style>