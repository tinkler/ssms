// Code generated by github.com/tinkler/mqttadmin; DO NOT EDIT.

import { _HttpClient } from '@delon/theme';

import { modelUrlPrefix } from './const';

export interface Role {
  id: number;

  name: string;

  saveRole(): Promise<void>;
}

export class Role {
  id: number = 0;

  name: string = '';

  constructor(private http: _HttpClient) {}

  saveRole(): Promise<void> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/role/role/save-role`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: {} }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.name = res.data.data['name'];

            resolve();
          } else {
            reject(res.message);
          }
        },
        error: err => {
          reject(err);
        }
      });
    });
  }
}
