const fs = require('fs');

const interestFactor = (photo1, photo2) => {
  const intersect = photo1.tags.filter(value => -1 !== photo2.tags.indexOf(value));
  const diff1 = photo1.tags.filter(value => -1 === photo2.tags.indexOf(value));
  const diff2 = photo2.tags.filter(value => -1 === photo1.tags.indexOf(value));
  return Math.min(diff1.length, intersect.length, diff2.length);
}

const fileName = process.argv[2];

const file = fs.readFileSync(fileName, { encoding: 'ascii' });

let data = file.split('\n');
data.shift();
data.pop();
let photos = data.map((s, i) => {
  return {
    i: String(i),
    o: s.split(' ')[0],
    tags: s.split(/\d /g)[1].split(' '),
  }
});

const hPhotos = photos.filter(photo => photo.o === 'H');
const vPhotos = photos.filter(photo => photo.o === 'V');
if (vPhotos.length % 2) vPhotos.pop();
let vSlides = [];
for (let k = 0; k < vPhotos.length; k += 2) {
  let vSlide = {
    i: `${vPhotos[k].i} ${vPhotos[k + 1].i}`,
    tags: [...new Set(vPhotos[k].tags.concat(vPhotos[k + 1].tags))],
  }
  vSlides.push(vSlide);
}
const slides = hPhotos.concat(vSlides);

let resultSlides = [slides.shift()];
while (slides.length) {
  let maxFactor = interestFactor(resultSlides[resultSlides.length - 1], slides[0]);
  let nextSlide = slides[0];
  for(const slide of slides) {
    const currentFactor = interestFactor(resultSlides[resultSlides.length - 1], slide);
    if (currentFactor > maxFactor) {
      maxFactor = currentFactor;
      nextSlide = slide;
    }
  }
  resultSlides.push(nextSlide);
  slides.splice(slides.indexOf(nextSlide), 1);
}

const result = resultSlides.reduce((res, slide) => `${res}${slide.i}\n`, `${resultSlides.length}\n`);

console.log(result);
