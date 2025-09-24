Build a go app that help me quick start a nextjs application

1. Init a nextjs project using command "`pnpx create-next-app@latest --yes [project-name]`"
2. delete unnecessary files
    1. `/public` folder
    2. `/app/favicon.ico`
    3. `README.md`
3. replace `/app/page.tsx` content with
    
    ```tsx
    export default function Home() {
      return <h1 className="text-3xl font-bold underline">Hello World</h1>;
    }
    ```
    
4. Install shadcn
    
    ```bash
    pnpm dlx shadcn@latest init -b neutral
    ```
    
5. do a git
    
    ```bash
    git add .
    git commit --amend --no-edit 
    ```
    

I want my binary file take a argument that is [project-name]