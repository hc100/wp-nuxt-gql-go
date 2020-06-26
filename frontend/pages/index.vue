<template>
  <v-app id="blog">
    <v-navigation-drawer v-model="drawer" app>
      <v-list dense>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Home</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <ArchiveCard />
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app color="indigo" dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>hacked 100%</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <PostPaging />
    </v-main>

    <v-footer color="indigo" app>
      <span class="white--text">&copy; 2019</span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import PostPaging from '~/components/PostPaging.vue'
import ArchiveCard from '~/components/ArchiveCard.vue'

@Component({
  components: { PostPaging, ArchiveCard },
})
export default class IndexPage extends Vue {
  private drawer = null
  private title = 'hacked 100%'

  head() {
    return {
      title: this.title,
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'Mac 1台で稼いで世界を旅しています',
        },
      ],
    }
  }
}
</script>
