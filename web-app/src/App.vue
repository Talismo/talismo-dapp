<template>
  <v-app>
    <app-bar />
    <v-main>
      <advertisement-top margin-top="2" margin-bottom="2"></advertisement-top>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import AppBar from "./components/navs/AppBar.vue";
import AdvertisementTop from "./components/advertisements/AdvertisementTop.vue";

export default {
  name: "App",
  components: {
    AppBar,
    AdvertisementTop,
  },
  data: () => ({
    //
  }),
};
</script>
