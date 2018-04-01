import { Component, OnInit } from '@angular/core';

import {Probe} from "../models/hardware";
import {TemperatureReading} from "../models/temperature";
import {BbqberryService} from '../bbqberry.service';
import {Observable} from "rxjs/Rx";


@Component({
  selector: 'app-temperature-indicator',
  templateUrl: './temperature-indicator.component.html',
  styleUrls: ['./temperature-indicator.component.css']
})
export class TemperatureIndicatorComponent implements OnInit {
  public probe: Probe;
  public reading: TemperatureReading;
  private pollingData: any;

  constructor(private bbqService: BbqberryService) { }

  ngOnInit() {
    this.bbqService.getHardwareConfig()
      .subscribe(hwCfg => this.probe = hwCfg.probes[0]);

    this.pollingData = Observable.interval(1000).startWith(0)
      .switchMap(() => this.bbqService.getTemperatureReading(0))
      .subscribe((readings) => {
        this.reading = readings[0];
      });
  }

  ngOnDestroy() {
    this.pollingData.unsubscribe();
  }

}
