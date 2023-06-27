import { NgModule, Optional, SkipSelf } from '@angular/core';

import { UserService } from './api/mqtt/user.service';
import { throwIfAlreadyLoaded } from './module-import-guard';

@NgModule({
  providers: [UserService]
})
export class CoreModule {
  constructor(@Optional() @SkipSelf() parentModule: CoreModule) {
    throwIfAlreadyLoaded(parentModule, 'CoreModule');
  }
}
