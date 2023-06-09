// Code generated by github.com/tinkler/mqttadmin; DO NOT EDIT.

import { _HttpClient } from '@delon/theme';

import { modelUrlPrefix } from './const';
import { Role } from './role.model';

export interface UserProfile {
  phoneNo: string;

  save(): Promise<void>;
}

export class UserProfile {
  phoneNo: string = '';

  constructor(private http: _HttpClient) {}

  save(): Promise<void> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/user_profile/save`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: {} }; message: string }) => {
          if (res.code === 0) {
            this.phoneNo = res.data.data['phone_no'];

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

export interface UserRole {
  id: number;

  user: User;

  role: Role;

  save(): Promise<void>;
}

export class UserRole {
  id: number = 0;

  user: User = new User(this.http);

  role: Role = new Role(this.http);

  constructor(private http: _HttpClient) {}

  save(): Promise<void> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/user_role/save`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: {} }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.user = res.data.data['user'];
            this.role = res.data.data['role'];

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

export interface Auth {
  /**
   * UUID
   */
  id: string;

  /**
   * UUID
   */
  deviceToken: string;

  username: string;

  password: string;

  token: string;

  signin(): Promise<Auth>;

  /**
   * QuickSignin quick signin with password
   */
  quickSignin(): Promise<void>;

  signup(): Promise<Auth>;
}

export class Auth {
  /**
   * UUID
   */
  id: string = '';

  /**
   * UUID
   */
  deviceToken: string = '';

  username: string = '';

  password: string = '';

  token: string = '';

  constructor(private http: _HttpClient) {}

  signin(): Promise<Auth> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/auth/signin`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: any }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.deviceToken = res.data.data['device_token'];
            this.username = res.data.data['username'];
            this.password = res.data.data['password'];
            this.token = res.data.data['token'];

            const resp: Auth = new Auth(this.http);
            resp.id = res.data.resp['id'];
            resp.deviceToken = res.data.resp['device_token'];
            resp.username = res.data.resp['username'];
            resp.password = res.data.resp['password'];
            resp.token = res.data.resp['token'];
            resolve(resp);
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

  /**
   * QuickSignin quick signin with password
   */
  quickSignin(): Promise<void> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/auth/quick-signin`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: {} }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.deviceToken = res.data.data['device_token'];
            this.username = res.data.data['username'];
            this.password = res.data.data['password'];
            this.token = res.data.data['token'];

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

  signup(): Promise<Auth> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/auth/signup`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: any }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.deviceToken = res.data.data['device_token'];
            this.username = res.data.data['username'];
            this.password = res.data.data['password'];
            this.token = res.data.data['token'];

            const resp: Auth = new Auth(this.http);
            resp.id = res.data.resp['id'];
            resp.deviceToken = res.data.resp['device_token'];
            resp.username = res.data.resp['username'];
            resp.password = res.data.resp['password'];
            resp.token = res.data.resp['token'];
            resolve(resp);
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

/**
 * User is the user model
 */
export interface User {
  /**
   * ID is the primary key
   */
  id: string;

  username: string;

  email: string;

  profiles: UserProfile[];

  /**
   * Save saves the user to the database
   */
  save(): Promise<void>;

  /**
   * AddRole adds a role to the user
   */
  addRole(role: Role): Promise<void>;
}

export class User {
  /**
   * ID is the primary key
   */
  id: string = '';

  username: string = '';

  email: string = '';

  profiles: UserProfile[] = [];

  constructor(private http: _HttpClient) {}

  /**
   * Save saves the user to the database
   */
  save(): Promise<void> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/user/save`, { data: this, args: {} }).subscribe({
        next: (res: { code: number; data: { data: any; resp: {} }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.username = res.data.data['username'];
            this.email = res.data.data['email'];
            this.profiles = res.data.data['profiles'];

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

  /**
   * AddRole adds a role to the user
   */
  addRole(role: Role): Promise<void> {
    return new Promise((resolve, reject) => {
      this.http.post(`${modelUrlPrefix}/user/user/add-role`, { data: this, args: { role: role } }).subscribe({
        next: (res: { code: number; data: { data: any; resp: {} }; message: string }) => {
          if (res.code === 0) {
            this.id = res.data.data['id'];
            this.username = res.data.data['username'];
            this.email = res.data.data['email'];
            this.profiles = res.data.data['profiles'];

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
