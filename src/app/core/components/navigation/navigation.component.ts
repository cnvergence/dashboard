import {Component, OnInit} from "@angular/core";
import { MdDialog } from '@angular/material';

import {Auth} from "../../services";
import {SidenavService} from "../sidenav/sidenav.service";
import {Store} from "@ngrx/store";
import * as fromRoot from "../../../redux/reducers/index";
import {Router} from '@angular/router';
import {environment} from "../../../../environments/environment";
import {AppConstants} from '../../../shared/constants/constants';
import { MobileNavigationComponent } from '../../../overlays';

@Component({
  selector: "kubermatic-navigation",
  templateUrl: "./navigation.component.html",
  styleUrls: ["./navigation.component.scss"]
})
export class NavigationComponent implements OnInit {

  public isScrolled: boolean = false;
  public environment: any = environment;

  // public userProfile: any;

  constructor(
    public auth: Auth, 
    private sidenavService: SidenavService, 
    private store: Store<fromRoot.State>,
    private router: Router,
    private dialog: MdDialog
  ) {}

  ngOnInit(): void {
    if (window.innerWidth < AppConstants.MOBILE_RESOLUTION_BREAKPOINT) {
      this.sidenavService.close();
    }
  }

  public logout() {
    this.router.navigate(['']);
    this.auth.logout();
  }

  public scrolledChanged(isScrolled) {
    this.isScrolled = isScrolled;
  }

  public toggleSidenav() {
    this.sidenavService
      .toggle()
      .then(() => { });
  }

  public onResize(event): void {
    if(event.target.innerWidth < AppConstants.MOBILE_RESOLUTION_BREAKPOINT) {
      this.sidenavService.close();
    }
  }

  public showMobileNav(): void {
    this.dialog.open(MobileNavigationComponent);
  }
}
