<template>
  <v-app-bar app color="light-blue darken-1" dark>
    <div class="d-flex align-center">
      <a href="/">
        <v-img
          alt="Talismo Logo"
          class="shrink mr-2"
          contain
          src="../../assets/logo-talismo.png"
          transition="scale-transition"
          width="50"
          to="home"
        />
      </a>
    </div>

    <v-spacer></v-spacer>

    <nft-search-button></nft-search-button>
    <utility-search-button></utility-search-button>
    <profile-search-button></profile-search-button>
  </v-app-bar>
</template>

<script>
import ProfileSearchButton from "./ProfileSearchButton.vue";
import UtilitySearchButton from "./UtilitySearchButton.vue";
import NftSearchButton from "./NftSearchButton copy.vue";

export default {
  components: {
    ProfileSearchButton,
    UtilitySearchButton,
    NftSearchButton,
  },
};
</script>

<style></style>
