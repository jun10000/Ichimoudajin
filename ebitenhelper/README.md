# Ebiten Helper（仮）

Ebitengineをより使いやすくするために開発されたライブラリです。  
主に下記の機能を追加します。  
- Tiled Map Editorにより作成されたマップデータの読み込み
- ウィジェットデータの読み込み（後述の"ウィジェットデータの仕様"に従う）
- キャラクターアニメーションの表示
- キャラクターの滑らかな動作
- コリジョン検出
- AI経路探索（A*アルゴリズム使用）

## ウィジェットデータの仕様（Ver.1）

XML形式で記述します。  

```xml
<?xml version="1.0" encoding="UTF-8"?>
<widget version="1">
    <!-- ここにウィジェットを配置していく -->
    <hbox bgcolor="#40000000">
        <text origin="0,0.5">テキストA</text>
        <text origin="0,0.5">テキストB</text>
    </hbox>
    <hbox origin="1,0" bgcolor="#40000000">
        <text origin="0,0.5">テキストC</text>
        <text origin="0,0.5">テキストD</text>
    </hbox>
    <vbox origin="0.5,0.5">
        <text origin="0.5,0">テキストE</text>
        <button origin="0.5,0" fgcolor="#ffff80">ボタンA</button>
    </vbox>
</widget>
```

一番外の要素は必ずwidget要素でなければいけません。  
その中にウィジェットを追加していき、レイアウトを構築していきます。  
ウィジェットはコンテナ要素、インライン要素に分けることができます。  

### コンテナ要素

子要素を持つことができる要素です。  

- widget  
  必須、1つだけ  
  要素を好きな位置に配置できます。  
- hbox  
  子要素を横方向に並べることができます。  
- vbox  
  子要素を縦方向に並べることができます。  

### インライン要素

子要素を持つことができない代わりに、様々な描画を行える要素です。  

- text  
  テキストを表示できます。  
- button  
  ボタンを表示できます。  

### 属性一覧

|属性|データ型|適用可能|継承|説明|
|---|-------|-------|---|---|
|version|Int|widget|-|ウィジェットデータのバージョン|
|name|String|全て|-|名前|
|origin|Vector|全て|-|配置推奨エリアのサイズ割合オフセット - 自要素のサイズ割合オフセット<br>例）左上揃え：(0,0)、右上揃え：(100,0)、右下揃え：(100,100)、水平垂直中央揃え：(50,50)|
|offset|Vector|全て|-|位置のオフセット|
|margin|Inset|全て|-|外側の余白|
|padding|Inset|全て|-|内側の余白|
|hide|Bool|全て|-|自要素以下の全要素を非表示にするか|
|bdwidth|Float|全て|-|枠線の太さ|
|bdcolor|Color|全て|-|枠線色|
|bgcolor|Color|全て|-|背景色|
|fgcolor|Color|全て|-|前景色|
|fontfiles|Strings|全て|〇|フォントファイルの相対パスリスト|
|fontsize|Float|全て|〇|フォントサイズ|

#### データ型一覧

複数の値はカンマ(,)で区切る  

|データ型|構造|説明|
|------|----|---|
|String|文字列|-|
|Bool|true/false|-|
|Int|整数|-|
|Float|小数点|基準（100）に対する割合を示す。<br>特に指定がない限り、画面サイズの高さを基準とする。|
|Color|16進カラーコード<br>（#RRGGBB/#AARRGGBB）|-|
|Strings|String x0~|-|
|Vector|Float x2|XY座標を示す。|
|Inset|Float x1/x2/x4|上下左右の各大きさを示す。<br>x1：（上下左右）、x2:（上下,左右）、x4:（上,右,下,左）|
