# What?

This is my personal CLI for working with Unreal Engine 5 on macOS. Currently my setup is using UE5.6 on macOS Tahoe 26.5.1. This is my personal CLI so it's not been tested outside the scope of my own setup.

I use it to do the following:
1. Compile and run my UE project
2. Generate compile_commands.json for use with clangd
3. Adding new UE5 C++ classes (actors, components, data, objects and plain C++ classes)

It's still under active development. There is no plan. I just add stuff to it when I need it. 

# The story so far

My old Windows game dev laptop was getting a little tired. It is a Lenovo Legion 7i. I do most of my game dev after hours, on weekends and holidays. So I want mobility. Although I was still able to use it with Unreal 5.4 it's battery life wasn't great. 

I got a Nov 2024 MacBook Pro from a tech start-up I was working at that went bust. It was just sitting on my desk gathering dust. So I decided to see if I could get my UE5 game project running on it. 

However, I ran into several issues once UE5.6 was running on it. The main one was that UE5.6 and Visual Studio ate up all the memory! There are also still some issue around file permissions I haven't solved yet. But I can live with that for now.

While I was doing game dev professionaly Visual Studio was definitely my preferred IDE. When not coding games I preferred to use old school editors like Emacs or Vim. When I started doing backend dev as my day job I moved to NeoVim and I'm totally in love with it. I already had NeoVim configured on the Mac I was trying to get my UE project running on. Many hours of tinkering went into getting clangd to work with UE. After I got that into a workable state (it's about 90% there) I found there were a few things missing from my workflow.

My first attempt was to implement these missing bits in NeoVim plugin. It worked fine, but I didn't want something tied to the editor. It also meant that I had to get more familiar with the NeoVim Lua API. I rather settled on using Go + Cobra to write a CLI. From a technical side it was perfect for what I wanted to build. Plus, I had to learn Go for my new job anyway, so here we are :)



