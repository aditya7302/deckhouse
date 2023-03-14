import type { Meta, Story } from "@storybook/vue3";

import BaseLayout from "@/components/layout/BaseLayout.vue";
import DeckhouseSettingsPage from "@/pages/DeckhouseSettingsPage.vue";

export default {
  title: "Deckhouse UI/Pages/DeckhouseSettings",
  component: DeckhouseSettingsPage,
  parameters: { layout: "fullscreen" },
} as Meta;

const Template: Story = (args, { loaded: { releases } }) => ({
  components: { DeckhouseSettingsPage, BaseLayout },

  setup() {
    return { args, releases: releases };
  },

  template: `
    <BaseLayout>
      <DeckhouseSettingsPage/>
    </BaseLayout>
  `,
});

export const Default = Template.bind({});