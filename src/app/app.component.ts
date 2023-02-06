import { Component } from '@angular/core';
import { URLList } from './constants';

interface Link {
  URL: string;
  Title: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'random-nytimes-recipe';
  public displayRecipeList: boolean = false;
  public recipeList: Link[] = [];

  constructor() {}

  ngOnInit() {
    console.log(URLList[Math.floor(Math.random() * URLList.length)]);
  }

  public generateRandomRecipeList() {
    this.recipeList = [];
    for (let i = 0; i < 5; i++) {
      let randomURL = this.getRandomRecipe();
      let randomRecipe: Link = {
        URL: randomURL,
        Title: this.getRecipeTitleFromURL(randomURL),
      };
      this.recipeList.push(randomRecipe);
    }
  }

  public goToRandomRecipe() {
    window.location.href = this.getRandomRecipe();
  }

  private getRecipeTitleFromURL(url: string): string {
    let title = url.replace('https://cooking.nytimes.com/recipes/', '');
    title = title.replace(/^\d+/g, '');
    title = this.toTitleCase(title);
    return title;
  }

  private toTitleCase(oldString: string): string {
    let newString = oldString
      .split('-')
      .map(function (word) {
        return word.charAt(0).toUpperCase() + word.slice(1);
      })
      .join(' ');

    return newString;
  }

  private getRandomRecipe() {
    return URLList[Math.floor(Math.random() * URLList.length)];
  }
}
