import React from "react";

const Home = ({ racks, role }) => {
  return (
    <div>
      <h1>ラック一覧</h1>
      <table border="1">
        <thead>
          <tr>
            <th>ラック番号</th>
            <th>高さ</th>
            <th>幅</th>
            <th>仕様</th>
            <th>奥行き</th>
          </tr>
        </thead>
        <tbody>
          {racks && racks.length > 0 ? (
            racks.map((rack) => (
              <tr key={rack.UnitNumber}>
                <td>{rack.UnitNumber}</td>
                <td>{rack.Height}</td>
                <td>{rack.Width}</td>
                <td>{rack.Specification}</td>
                <td>{rack.Depth}</td>
                <td>
                  <a
                    href={`/rack/${rack.UnitNumber}/`}
                    onClick={(e) => {
                      if (
                        role !== "rack" &&
                        role !== "nw" &&
                        role !== "server" &&
                        role !== "admin" &&
                        role !== "guest"
                      ) {
                        e.preventDefault();
                      }
                    }}
                    style={{
                      color:
                        role !== "rack" && role !== "admin"
                          ? "gray"
                          : "inherit",
                    }}
                  >
                    ラックの詳細
                  </a>
                  <a
                    href={`/edit/${rack.UnitNumber}/`}
                    onClick={(e) => {
                      if (role !== "rack" && role !== "admin") {
                        e.preventDefault();
                      }
                    }}
                    style={{
                      color:
                        role !== "rack" && role !== "admin"
                          ? "gray"
                          : "inherit",
                    }}
                  >
                    情報の編集
                  </a>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="5">該当するレコードが見つかりません。</td>
            </tr>
          )}
        </tbody>
      </table>

      <form action="/add" method="post">
        <input type="text" name="unitNumber" placeholder="ラック番号" />
        <input type="text" name="height" placeholder="高さ" />
        <input type="text" name="width" placeholder="幅" />
        <input type="text" name="specification" placeholder="仕様" />
        <input type="text" name="depth" placeholder="奥行き" />
        <input type="submit" value="追加" />
      </form>

      <form action="/delete" method="post">
        <input type="text" name="unitNumber" placeholder="ラック番号" />
        <input type="submit" value="削除" />
      </form>

      <form action="/search" method="get">
        <input type="text" name="unitNumber" placeholder="ラック番号" />
        <input type="submit" value="検索" />
      </form>
    </div>
  );
};

export default Home;
