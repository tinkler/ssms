import { Environment } from '@delon/theme';

export const environment = {
  production: true,
  useHash: true,
  api: {
    baseUrl: 'http://localhost:11080/',
    refreshTokenEnabled: false,
    refreshTokenType: 'auth-refresh'
  }
} as Environment;
