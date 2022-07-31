<template>
    <section class="section">
        <dev class="container">
            <div class="block">
                <div class="block" vfor="item in items">
                    <AppVideo
                        :item="item"
                        :video-id="item.id"
                    />
                </div>
            </div>
        </dev>
    </section>
</template>



<script>
import ROUTES from '~/routes/api'
import AppVideo from "~/components/AppVideo"

export default {
    components:{
        AppVideo
    },

    computed:{
        items(){
            return this.$store.getters.getFavoriteVideos
        }
    },

    async fetch({store}){
        const payload = {
            uri:ROUTES.GET.FAVORITE,
        }

        await store.dispatch('fetchFavoriteVideos', payload)
    }
}
</script>