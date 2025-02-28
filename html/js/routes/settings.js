// Copyright 2020 Security Onion Solutions. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

routes.push({ path: '/settings', name: 'settings', component: {
  template: '#page-settings',
  data() { return {
    i18n: this.$root.i18n,
    showSettingsForm: false,
    showPassword: false,
    usingDefaults: false,
    form: {
      valid: false,
      email: null,
      password: null,
      csrfToken: null,
    },
    rules: {
      required: value => !!value || this.$root.i18n.required,
      matches: value => (!!value && value == this.form.password) || this.$root.i18n.passwordMustMatch,
    },
    authSettingsUrl: null,
  }},
  mounted() {
    if (!this.$root.getAuthFlowId()) {
      this.reloadSettings();
    } else {
      this.showSettingsForm = true;
      this.authSettingsUrl = this.$root.authUrl + 'settings/methods/password' + location.search;
      this.loadData()
    }
    this.usingDefaults = localStorage.length == 0;
  },
  watch: {
  },
  methods: {
    reloadSettings() {
      location.pathname = this.$root.settingsUrl;
    },
    resetDefaults() {
      localStorage.clear();
      this.usingDefaults = true;
    },
    async loadData() {
      try {
        const response = await this.$root.authApi.get('settings/flows?id=' + this.$root.getAuthFlowId());
        this.form.csrfToken = response.data.methods.password.config.fields.find(item => item.name == 'csrf_token').value;
        var errors = [];
        response.data.methods.password.config.fields.forEach(function(value, index, array) {
          if (value.messages) {
            value.messages.forEach(function(err, idx, errArray) {
              errors.push(err.text);
            });
          }
        });
        if (errors.length > 0) {
          this.$root.showWarning(this.i18n.settingsInvalid + errors.join("\n"));
        } else if (response.data.state == "success") {
          this.$root.showInfo(this.i18n.settingsSaved);
        }
      } catch (error) {
        if (error != null && error.response != null && error.response.status == 410) {
          this.reloadSettings();
        } else {
          this.$root.showError(error);
        }
      }
    },
  }
}});
