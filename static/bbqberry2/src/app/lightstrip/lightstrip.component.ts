import { Component, OnInit } from '@angular/core';
import {LightStrip} from "../models/lightstrip";
import {BbqberryService} from "../bbqberry.service";
import {Observable} from "rxjs/Rx";

@Component({
  selector: 'app-lightstrip',
  templateUrl: './lightstrip.component.html',
  styleUrls: ['./lightstrip.component.css']
})
export class LightstripComponent implements OnInit {
  private strip: LightStrip;
  private pollingInterval: number;
  private pollingData: any;

  constructor(private bbqService: BbqberryService) { }

  ngOnInit() {
    // this.setupPolling(100);
  }

  ngOnDestroy() {
  }

  private stopPolling() {
    this.pollingInterval = 0;
    this.pollingData.unsubscribe();
  }

  private setupPolling(interval: number) {
    // console.log(`SETTING UP POLLING with ${interval}`);
    this.pollingData = Observable.interval(interval).startWith(0)
      .switchMap(() => this.bbqService.getGrillLightStrip())
      .subscribe((strip) => {
        if ( strip == null ) {
          this.stopPolling()
        }
        else {
          let reset = this.strip != null && this.strip.interval / 1000 != interval;
          this.strip = strip;

          if ( reset ) {
            let interval = this.strip.interval / 1000;
            this.stopPolling();
            this.setupPolling(interval);
          }
        }
      });
  }
}
