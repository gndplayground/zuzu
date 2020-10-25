# zuzu

`zuzu` is a small utility that helps generate files from a template. Just something like angular-cli `ng generate component` but not focused on any framework.

## Install

```
npm install --save-dev zuzu
```

```
yarn add --dev zuzu
```

## Usage

```
npx zuzu name
```

By the default the CLI will look for the `template` folder in the current calling directory and copy the content of that folder to the name was inputted.

```markdown
node_modules
template
├──── file.js
name
├──── file.js
```

### -base-template

We can change the default template location by using `-base-template` arg.

```
npx zuzu -base-template=custom/location name
```

### -t

If the template location has multiple directories and we want to select a specific folder we can use `-t` arg.

```
npx zuzu -t=cat name
```

```markdown
node_modules
template
├──── cat
├──── cat.js
├──── dog
├──── dog.js
name
├──── cat.js
```

### -no-dir

If we do not want to create the directory in the current calling location we can use `-no-dir` arg.

```
npx zuzu -no-dir name
```

```markdown
node_modules
template
├──── file.js
file.js
```

### -dir

We can change the output directory location using`-dir` arg

```
npx zuzu -dir=hello name
```

```markdown
node_modules
template
├──── file.js
hello
├──── name
├──── file.js
```

## Variable

The CLI support name variable.

```
npx zuzu hello
```

In the above cmd, the name variable is "hello". We can replace text content or file name in the selected template folder by using the name variable.

- {{name}} => replace to current name.
- {{nameCamel}} => replace to lower camel case
- {{NameCamel}} => replace to camel case
- {{nameKebab}} => replace to kebab case
- {{NameKebab}} => replace to screaming kebab case.
- {{Name}} => replace to title case.
- {{NAME}} => replace to upper case.

### Example

```
npx zuzu hello-world
```

```markdown
node_modules
template
├──── {{NameCamel}}.js
hello-world
├──── HelloWorld.js
```

Content {{NameCamel}}.js file

```
import React, {FunctionComponent} from 'react';

const {{NameCamel}}: FunctionComponent = () => {
    return (
        <h1>Hello!</h1>
    );
};

export default {{NameCamel}};
```

Content HelloWorld.js file

```
import React, {FunctionComponent} from 'react';

const HelloWorld: FunctionComponent = () => {
    return (
        <h1>Hello!</h1>
    );
};

export default HelloWorld;
```
