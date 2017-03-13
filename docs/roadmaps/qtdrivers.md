# Roadmap

1. Get QT Webview working with Greeter.
This is just to make sure we are good with building a very basic example.
- Not need to do anything else.
- Intent is so that the exact same code for the Web target, also works with the QT WebView. If we find out that there are slight changes lets review if we find them.
- Base example is here: https://github.com/therecipe/qt/tree/master/internal/examples/qml/webview

2. Get QT Greeter working on Win 10, OSX, Ubuntu & Android, IOS.
This is to make sure we can build on all targets.

Now at his point there are a few things we can focus on, but i think its more impotant to state what we dont want to do that will create lots of work and confusion.
We want commonise on a standard way that works on all targets (web, Desktop, Mobile) so that its easy to have "write once, run everywhere".
What we dont want to do:
- Support native menus on the Desktop Qt versions. Mobile and web does not have this, so lets not do it on the Desktop versions. Its far better to build all menu 7 nav stuff in the actual HTML UI.

3. Back button, with pop and push stack concept.
This makes navigation work on mobile, desktop & web within a single app.
This is important because browser, Desktop and mobile have slight different concept of "back".

- The way the QT Draw example works in the therecipt/QT example is perfect i feel.
  - https://github.com/therecipe/qt/blob/master/internal/examples/qml/drawer_nav_x/navigation/HomeNavigation.qml#L26
  - https://github.com/therecipe/qt/blob/master/internal/examples/qml/gallery/pages/StackViewPage.qml
  - https://github.com/therecipe/qt/blob/master/internal/examples/qml/gallery/pages/SwipeViewPage.qml
    - Not sure we need swipe yet, if at all.

4. Intents.
This makes navigation work on mobile, desktop & web, between apps.

Now we can start to look at how the Concept of larger applications can work in terms of "windows".
The best way to explan this is with a simple example. Suppose your making an application that does inventory management for warehouses.
You have a main Window that is where users can add, remove, edit inventory.
You want all users to be able to chat about the inventory over a Window that works like Google Hangout.
You will probably introduce more things to the Application later too like users management, and who knows what else.

Now on Web, Desktop and Mobiles these things need to behave a bit differently, but not too much If we follow the concept above:
- Web
  - Inventory is a Tab.
  - Hangout is a Tab.
- Desktop
  - Inventory is a Window.
  - Hangout is a Window.
- Mobile
  - Inventory is an App.
  - Hangout is an App.

 Now the way that everything is done on Mobile is via Intents, and for web there is "web-intents". This is a really smart way of allowing Apps to work with each other in such as way that they do not knwo about each other.
 Its a bit like how Modules work in a CMS.
 Basically the Inventory App, exposes a "share" intent, and the Hangout App subscribes to "shares".
 This is really exactly like pub sub, but for applicatiosn themsleves running on computer.

 http://webintents.org/
 - This gives great examples.
 https://en.wikipedia.org/wiki/Web_Intents
 -  Now web intents never took off, but we can use if for GU system for sure.
 - Android has intents because they work so well, but the rest of the browsers never embraced it.
 - IOS also has a simialr thing, that we can tap into.
 All we need to do is make an interface that all GU code users, and a slightly different implementation for each Target.
 On Andrpid and IOS the are expressed in the manifest file that is standard for apps.
 We can generate this, this is why intents are loosely coupled.


Now this is enough. Lots of other things for later are in the gu-x repo.

Gu Modules can now be looked at, but no point yet