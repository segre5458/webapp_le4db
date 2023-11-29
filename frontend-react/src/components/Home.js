import React from "react";

const Home = ({ racks, role }) => {
  return (
    <div>
      <h1>���b�N�ꗗ</h1>
      <table border="1">
        <thead>
          <tr>
            <th>���b�N�ԍ�</th>
            <th>����</th>
            <th>��</th>
            <th>�d�l</th>
            <th>���s��</th>
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
                    ���b�N�̏ڍ�
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
                    ���̕ҏW
                  </a>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="5">�Y�����郌�R�[�h��������܂���B</td>
            </tr>
          )}
        </tbody>
      </table>

      <form action="/add" method="post">
        <input type="text" name="unitNumber" placeholder="���b�N�ԍ�" />
        <input type="text" name="height" placeholder="����" />
        <input type="text" name="width" placeholder="��" />
        <input type="text" name="specification" placeholder="�d�l" />
        <input type="text" name="depth" placeholder="���s��" />
        <input type="submit" value="�ǉ�" />
      </form>

      <form action="/delete" method="post">
        <input type="text" name="unitNumber" placeholder="���b�N�ԍ�" />
        <input type="submit" value="�폜" />
      </form>

      <form action="/search" method="get">
        <input type="text" name="unitNumber" placeholder="���b�N�ԍ�" />
        <input type="submit" value="����" />
      </form>
    </div>
  );
};

export default Home;
