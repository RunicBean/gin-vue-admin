<script setup>
import MainImages from "./layoutOne/MainImages.vue";
import DetailedItems from "./layoutOne/DetailedItems.vue";
import {onMounted, ref} from "vue";
import {findCompanyByPath} from "@/api/company";
import {useRoute, useRouter} from "vue-router";

const route = useRoute()
const router = useRouter()
const companyName = ref("")
onMounted(() => {
    findCompanyByPath({path: route.params.companyPath})
        .then((res) => {
            if (res.code !== 0) {
                router.push({ path: '/layout/404' })
            }
            companyName.value = res.data.recompany.name
            document.title = companyName.value
        })
        .catch((err) => {
            router.push({ path: '/layout/404' })
        })
})
</script>

<template>
    <div class="body">
        <MainImages :company-name="companyName" />
        <DetailedItems :company-name="companyName" />
    </div>
</template>

<style scoped>
div {
    background-color: #fff;
}
</style>
