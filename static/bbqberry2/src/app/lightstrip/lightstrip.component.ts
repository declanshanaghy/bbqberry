import { Component, OnInit } from '@angular/core';
import {LightStrip} from "../models/lightstrip";
import {BbqberryService} from "../bbqberry.service";
import {Observable} from "rxjs/Observable";

@Component({
  selector: 'app-lightstrip',
  templateUrl: './lightstrip.component.html',
  styleUrls: ['./lightstrip.component.css']
})
export class LightstripComponent implements OnInit {
  private strip: LightStrip;
  private pollingData: any;

  constructor(private bbqService: BbqberryService) { }

  ngOnInit() {
    this.pollingData = Observable.interval(1).startWith(0)
      .switchMap(() => this.bbqService.getGrillLightStrip())
      .subscribe((strip) => {
        this.strip = strip;
      });

  }

  ngOnDestroy() {
    this.pollingData.unsubscribe();
  }

}
