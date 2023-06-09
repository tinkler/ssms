// Code generated by github.com/tinkler/mqttadmin; DO NOT EDIT.
import { Injectable } from '@angular/core';
import { _HttpClient } from '@delon/theme';

import { Role } from './role.model';
import { RoleService } from './role.service';
import { UserProfile, UserRole, Auth, User } from './user.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  constructor(private http: _HttpClient, private roleSrv: RoleService) {}

  newUserProfile(): UserProfile {
    return new UserProfile(this.http);
  }
  newUserRole(): UserRole {
    return new UserRole(this.http);
  }
  newAuth(): Auth {
    return new Auth(this.http);
  }
  /**
   * User is the user model
   */
  newUser(): User {
    return new User(this.http);
  }
}
