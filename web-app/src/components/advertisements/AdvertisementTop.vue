<template>
  <v-container class="my-2">
    <v-row align="center" justify="center">
      <v-col cols="10">
        <v-card elevation="1">
          <v-card-text> Announcement </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn plain> read more </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  props: ["margin-top", "margin-bottom"],
};
</script>

<style></style>
