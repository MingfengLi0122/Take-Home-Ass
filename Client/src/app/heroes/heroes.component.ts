import { Component, OnInit } from '@angular/core';
import { Hero } from '../hero';
import { HeroService } from '../hero.service';

// decorator
@Component({
  //  the component's CSS element selector
  selector: 'app-heroes',
  // the location of the component's template file.
  templateUrl: './heroes.component.html',
  //  the location of the component's private CSS styles.
  styleUrls: ['./heroes.component.css']
})
export class HeroesComponent implements OnInit {
  heroes: Hero[] = [];
  // ? marks member as being optional in the interface
  // selectedHero?: Hero;
  
  constructor(private heroService: HeroService) { }

  // calls it after createing component, great for initialize logic
  ngOnInit(): void {
    this.getHeroes();
  }
  
  // onSelect(hero: Hero): void {
  //   this.selectedHero = hero;
  //   this.messageService.add(`HeroesComponent: Selected hero id=${hero.id}`);
  // }

  getHeroes(): void {
    this.heroService.getHeroes()
      .subscribe(heroes => this.heroes = heroes);
  }

  add(name: string): void {
    name = name.trim();
    if (!name) { return; }
    this.heroService.addHero({ name } as Hero)
      .subscribe(hero => {
        this.heroes.push(hero);
      });
  }
  
  delete(hero: Hero): void {
    this.heroes = this.heroes.filter(h => h!==hero);
    this.heroService.deleteHero(hero.id).subscribe();
  }
}
